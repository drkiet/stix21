package stix21

type Binary byte
type Hashes map[string]interface{}
type Hex string
type Boolean bool
type Float float32
type Identifier string
type OpenVocab string

const DEFAULT_VERSION = "2.1"

// src: SDO/SCO target: SDO/SCO
const DerivedFromRelationship = "derived-from"
const DuplicateOfRelationship = "duplicate-of"
const RelatedToRelationship = "related-to"

// Reserved Name: RN
const SeverityRN = "severity"
const UserNamesRN = "usernames"
const PhoneNumbersRN = "phone_numbers"

// Reserved Object Type Name
const IncidentROTN = "incident"
const ActionROTN = "action"