package secrets

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
)

const (
	// SimpleType secrets are basic string secrets.
	SimpleType = "simple"

	// VersionedType secrets are secrets that can be rotated gracefully.
	VersionedType = "versioned"

	// CredentialType secrets are username/password pairs as a single secret
	// in vault.
	CredentialType = "credential"
)

// A Secret is the base type of secrets.
type Secret []byte

// IsEmpty returns true if the secret is empty.
func (s Secret) IsEmpty() bool {
	return len(s) == 0
}

// CSIFile represents the raw parsed object of a file made by the Vault CSI provider
type CSIFile struct {
	Secret GenericSecret `json:"data"`
}

// Secrets allows to access secrets based on their different type.
type Secrets struct {
	simpleSecrets     map[string]SimpleSecret
	versionedSecrets  map[string]VersionedSecret
	credentialSecrets map[string]CredentialSecret
	vault             Vault
}

// GetSimpleSecret fetches a simple secret or error if the key is not present.
func (s *Secrets) GetSimpleSecret(path string) (SimpleSecret, error) {
	if path == "" {
		return SimpleSecret{}, ErrEmptySecretKey
	}
	secret, ok := s.simpleSecrets[path]
	if !ok {
		return secret, SecretNotFoundError(path)
	}

	return secret, nil
}

// GetVersionedSecret fetches a versioned secret or error if the key is not present.
func (s *Secrets) GetVersionedSecret(path string) (VersionedSecret, error) {
	if path == "" {
		return VersionedSecret{}, ErrEmptySecretKey
	}
	secret, ok := s.versionedSecrets[path]
	if !ok {
		return secret, SecretNotFoundError(path)
	}

	return secret, nil
}

// GetCredentialSecret fetches a credential secret or error if the key is not
// present.
func (s *Secrets) GetCredentialSecret(path string) (CredentialSecret, error) {
	if path == "" {
		return CredentialSecret{}, ErrEmptySecretKey
	}
	secret, ok := s.credentialSecrets[path]
	if !ok {
		return secret, SecretNotFoundError(path)
	}

	return secret, nil
}

// SimpleSecret represent basic secrets.
type SimpleSecret struct {
	Value Secret
}

// Returns a new instance of SimpleSecret based on a
// GenericSecret from Document. If there is an encoding specified the
// raw secret will be decoded prior.
func newSimpleSecret(secret *GenericSecret) (SimpleSecret, error) {
	var result SimpleSecret
	value, err := secret.Encoding.decodeValue(secret.Value)
	if err != nil {
		return result, err
	}
	return SimpleSecret{
		Value: value,
	}, nil
}

// AsVersioned returns the SimpleSecret as a VersionedSecret.
//
// The Value of the SimpleSecret will be set as the Current value on the
// VersionedSecret.
func (s SimpleSecret) AsVersioned() VersionedSecret {
	return VersionedSecret{Current: s.Value}
}

// VersionedSecret represent secrets like signing keys that can be rotated
// gracefully.
//
// The current property contains the active version of a secret. This should be
// used for any actions that generate new cryptographic data (e.g. signing a
// token).
//
// The previous and next fields contain old and not-yet-active versions of the
// secret respectively. These MAY be used by applications to give a grace
// period for cryptographic tokens generated during a rotation, but SHOULD NOT
// be used to generate new cryptographic tokens.
type VersionedSecret struct {
	Current  Secret
	Previous Secret
	Next     Secret
}

// Returns a new instance of VersionedSecret based on a
// GenericSecret from Document. If there is an encoding specified the
// raw secrets will be decoded prior.
func newVersionedSecret(secret *GenericSecret) (VersionedSecret, error) {
	var result VersionedSecret

	current := secret.Current
	previous := secret.Previous
	next := secret.Next

	currentSecret, err := secret.Encoding.decodeValue(current)
	if err != nil {
		return result, err
	}
	previousSecret, err := secret.Encoding.decodeValue(previous)
	if err != nil {
		return result, err
	}
	nextSecret, err := secret.Encoding.decodeValue(next)
	if err != nil {
		return result, err
	}
	return VersionedSecret{
		Current:  currentSecret,
		Previous: previousSecret,
		Next:     nextSecret,
	}, nil
}

// GetAll returns all versions that are not empty in the following order:
// current, previous, next.
func (v *VersionedSecret) GetAll() []Secret {
	allVersions := []Secret{v.Current}
	if !v.Previous.IsEmpty() {
		allVersions = append(allVersions, v.Previous)
	}
	if !v.Next.IsEmpty() {
		allVersions = append(allVersions, v.Next)
	}
	return allVersions
}

// CredentialSecret represent represent username/password pairs as a single
// secret in vault. Note that usernames are not generally considered secret,
// but they are tied to passwords.
type CredentialSecret struct {
	Username string
	Password string
}

// NewCredentialSecret returns a new instance of CredentialSecret based on a
// GenericSecret from Document.
func newCredentialSecret(secret *GenericSecret) (CredentialSecret, error) {
	return CredentialSecret{
		Username: secret.Username,
		Password: secret.Password,
	}, nil
}

// Document represents the raw parsed entity of a Secrets JSON and is
// not meant to be used other than instantiating Secrets.
type Document struct {
	Secrets map[string]GenericSecret `json:"secrets"`
	Vault   Vault                    `json:"vault"`
}

// Validate checks the Document for any errors that violate the Baseplate
// specification.
//
// When this function returns a non-nil error, the error is either a
// TooManyFieldsError, or a BatchError containing multiple TooManyFieldsError.
func (s *Document) Validate() error {
	errs := make([]error, 0, len(s.Secrets))
	for key, value := range s.Secrets {
		if value.Type == SimpleType && notOnlySimpleSecret(value) {
			errs = append(errs, TooManyFieldsError{
				SecretType: SimpleType,
				Key:        key,
			})
		}
		if value.Type == VersionedType && notOnlyVersionedSecret(value) {
			errs = append(errs, TooManyFieldsError{
				SecretType: VersionedType,
				Key:        key,
			})
		}
		if value.Type == CredentialType && notOnlyCredentialSecret(value) {
			errs = append(errs, TooManyFieldsError{
				SecretType: CredentialType,
				Key:        key,
			})
		}
	}
	return errors.Join(errs...)
}

func notOnlySimpleSecret(secret GenericSecret) bool {
	return secret.Current != "" || secret.Previous != "" || secret.Next != "" || secret.Username != "" || secret.Password != ""
}

func notOnlyVersionedSecret(secret GenericSecret) bool {
	return secret.Value != "" || secret.Username != "" || secret.Password != ""
}

func notOnlyCredentialSecret(secret GenericSecret) bool {
	return secret.Value != "" || secret.Current != "" || secret.Previous != "" || secret.Next != ""
}

// GenericSecret is a placeholder to fit all types of secrets when parsing the
// Secret JSON before processing them into their more typed equivalents.
type GenericSecret struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Encoding Encoding `json:"encoding"`

	Current  string `json:"current"`
	Previous string `json:"previous"`
	Next     string `json:"next"`

	Username string `json:"username"`
	Password string `json:"password"`
}

// Vault provides authentication credentials so that applications can directly
// connect to Vault for more complicated use cases.
type Vault struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// NewSecrets parses and validates the secret JSON provided by the reader.
func NewSecrets(r io.Reader) (*Secrets, error) {
	var secretsDocument Document
	err := json.NewDecoder(r).Decode(&secretsDocument)
	if err != nil {
		return nil, err
	}

	return secretsValidate(secretsDocument)
}

// FromDir parses a directory and returns its secrets
func FromDir(dir fs.FS) (*Secrets, error) {
	secretsDocument, err := walkCSIDirectory(dir)
	if err != nil {
		return nil, err
	}
	return secretsValidate(secretsDocument)

}

func secretsValidate(secretsDocument Document) (*Secrets, error) {
	secrets := &Secrets{
		simpleSecrets:     make(map[string]SimpleSecret),
		versionedSecrets:  make(map[string]VersionedSecret),
		credentialSecrets: make(map[string]CredentialSecret),
		vault:             secretsDocument.Vault,
	}
	if err := secretsDocument.Validate(); err != nil {
		return nil, err
	}
	for key, secret := range secretsDocument.Secrets {
		switch secret.Type {
		case "simple":
			simple, err := newSimpleSecret(&secret)
			if err != nil {
				return nil, err
			}
			secrets.simpleSecrets[key] = simple
		case "versioned":
			versioned, err := newVersionedSecret(&secret)
			if err != nil {
				return nil, err
			}
			secrets.versionedSecrets[key] = versioned
		case "credential":
			credential, err := newCredentialSecret(&secret)
			if err != nil {
				return nil, err
			}
			secrets.credentialSecrets[key] = credential
		default:
			return nil, fmt.Errorf(
				"secrets.NewSecrets: encountered unknown secret type %q for secret %q",
				secret.Type,
				key,
			)
		}
	}
	return secrets, nil
}

type notCSIError struct {
	err error
}

func (e notCSIError) Error() string {
	return fmt.Sprintf("configured directory does not appear to be the root of a Vault CSI mount point: %v", e.err)
}

func (e notCSIError) Unwrap() error {
	return e.err
}

// walkCSIDirectory parses a directory for vault secrets and merges them into one object
func walkCSIDirectory(dir fs.FS) (Document, error) {
	const (
		// This is where k8s actually writes the content,
		// ref: https://pkg.go.dev/sigs.k8s.io/secrets-store-csi-driver/pkg/util/fileutil#AtomicWriter.Write
		k8sSubdirectory = "..data"
	)
	secretsDocument := Document{
		Secrets: make(map[string]GenericSecret),
	}
	err := fs.WalkDir(
		dir,
		k8sSubdirectory,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				if path == k8sSubdirectory && errors.Is(err, fs.ErrNotExist) {
					return notCSIError{err: err}
				}
				return err
			}
			if d.IsDir() {
				return nil
			}
			relPath, err := filepath.Rel(k8sSubdirectory, path)
			if err != nil {
				// Should not happen as this means path is not under k8sSubdirectory,
				// but just in case
				return nil
			}
			file, err := dir.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			var secretFile CSIFile
			err = json.NewDecoder(file).Decode(&secretFile)
			if err != nil {
				return fmt.Errorf("decoding %q: %w", path, err)
			}
			secretsDocument.Secrets[relPath] = secretFile.Secret
			return nil
		},
	)
	if err != nil {
		return Document{}, fmt.Errorf("secrets.walkCSIDirectory: %w", err)
	}
	return secretsDocument, nil
}
