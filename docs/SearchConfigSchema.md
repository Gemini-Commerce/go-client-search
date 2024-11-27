# # SearchConfigSchema
contains index configurations fields

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index**| **string** | index name, to create or where update configuration  |
**DocumentKeyName**| **string** | field which value will be used as document id [#DOCGENBUG REQUIRED FIELD]  | [optional]
**Attributes**| [**[]SearchConfigSchemaAttribute**](SearchConfigSchemaAttribute.md) | fields that can be stored into index and later retrieved  |
**Searchables**| [**[]SearchConfigSchemaSearchable**](SearchConfigSchemaSearchable.md) | fields that can be used for fulltext searches  | [optional]
**Facets**| [**[]SearchConfigSchemaFacet**](SearchConfigSchemaFacet.md) | fields that can be used for aggregations  | [optional]
**Filters**| [**[]SearchConfigSchemaFilter**](SearchConfigSchemaFilter.md) | fields that can be used for filtering  | [optional]
**Sortables**| [**[]SearchConfigSchemaSortable**](SearchConfigSchemaSortable.md) | fields that can be used for sorting  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

