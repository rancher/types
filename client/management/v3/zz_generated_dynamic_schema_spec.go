package client

const (
	DynamicSchemaSpecType                      = "dynamicSchemaSpec"
	DynamicSchemaSpecFieldCollectionActions    = "collectionActions"
	DynamicSchemaSpecFieldCollectionFields     = "collectionFields"
	DynamicSchemaSpecFieldCollectionFilters    = "collectionFilters"
	DynamicSchemaSpecFieldCollectionMethods    = "collectionMethods"
	DynamicSchemaSpecFieldDynamicSchemaVersion = "dynamicSchemaVersion"
	DynamicSchemaSpecFieldEmbed                = "embed"
	DynamicSchemaSpecFieldEmbedType            = "embedType"
	DynamicSchemaSpecFieldIncludeableLinks     = "includeableLinks"
	DynamicSchemaSpecFieldPluralName           = "pluralName"
	DynamicSchemaSpecFieldResourceActions      = "resourceActions"
	DynamicSchemaSpecFieldResourceFields       = "resourceFields"
	DynamicSchemaSpecFieldResourceMethods      = "resourceMethods"
	DynamicSchemaSpecFieldSchemaName           = "schemaName"
)

type DynamicSchemaSpec struct {
	CollectionActions    map[string]Action `json:"collectionActions,omitempty" yaml:"collection_actions,omitempty"`
	CollectionFields     map[string]Field  `json:"collectionFields,omitempty" yaml:"collection_fields,omitempty"`
	CollectionFilters    map[string]Filter `json:"collectionFilters,omitempty" yaml:"collection_filters,omitempty"`
	CollectionMethods    []string          `json:"collectionMethods,omitempty" yaml:"collection_methods,omitempty"`
	DynamicSchemaVersion string            `json:"dynamicSchemaVersion,omitempty" yaml:"dynamic_schema_version,omitempty"`
	Embed                bool              `json:"embed,omitempty" yaml:"embed,omitempty"`
	EmbedType            string            `json:"embedType,omitempty" yaml:"embed_type,omitempty"`
	IncludeableLinks     []string          `json:"includeableLinks,omitempty" yaml:"includeable_links,omitempty"`
	PluralName           string            `json:"pluralName,omitempty" yaml:"plural_name,omitempty"`
	ResourceActions      map[string]Action `json:"resourceActions,omitempty" yaml:"resource_actions,omitempty"`
	ResourceFields       map[string]Field  `json:"resourceFields,omitempty" yaml:"resource_fields,omitempty"`
	ResourceMethods      []string          `json:"resourceMethods,omitempty" yaml:"resource_methods,omitempty"`
	SchemaName           string            `json:"schemaName,omitempty" yaml:"schema_name,omitempty"`
}
