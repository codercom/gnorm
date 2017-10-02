package cli

// Config holds the schema that is expected to exist in the gnorm.toml file.
type Config struct {
	// ConnStr is the connection string for the database.  Environment variables
	// in $FOO form will be expanded.
	ConnStr string

	// The type of DB you're connecting to.  Currently the possible values are
	// "postgres" or "mysql".
	DBType string

	// Schemas holds the names of schemas to generate code for.
	Schemas []string

	// IncludeTables is a whitelist of tables to generate data for. Tables not
	// in this list will not be included in data geenrated by gnorm. You cannot
	// set IncludeTables if ExcludeTables is set.  By default, tables will be
	// included in all schemas.  To specify tables for a specific schema only,
	// use the schema.tablenmae format.
	IncludeTables []string

	// ExcludeTables is a blacklist of tables to ignore while generating data.
	// All tables in a schema that are not in this list will be used for
	// generation. You cannot set ExcludeTables if IncludeTables is set.  By
	// default, tables will be excluded from all schemas.  To specify tables for
	// a specific schema only, use the schema.tablenmae format.
	ExcludeTables []string

	// PostRun is a command with arguments that is run after each file is
	// generated by GNORM.  It is generally used to reformat the file, but it
	// can be for any use. Environment variables will be expanded, and the
	// special $GNORMFILE environment variable may be used, which will expand to
	// the name of the file that was just generated.
	PostRun []string

	// NameConversion defines how the DBName of tables, schemas, and enums are
	// converted into their Name value.  This is a template that may use all the
	// regular functions.  The "." value is the DB name of the item. Thus, to
	// make an item's Name the same as its DBName, you'd use a template of
	// "{{.}}". To make the Name the PascalCase version of DBName, you'd use
	// "{{pascal .}}".
	NameConversion string

	// TablePaths is a set of "output-path" = "template-path" pairs that tells
	// Gnorm how to render and output its table info.  Each template will be
	// rendered with each table in turn and written out to the given output
	// path. If no pairs are specified, tables will not be rendered.
	//
	// The table path may be a template, in which case the values .Schema and
	// .Table may be referenced, containing the name of the current schema and
	// table being rendered.  For example,
	// "{{.Schema}}/{{.Table}}/{{.Table}}.go" = "tables.gotmpl" would render
	// tables.gotmpl template with data from the the "public.users" table to
	// ./public/users/users.go.
	TablePaths map[string]string

	// SchemaPaths is a set of "output-path" = "template-path" pairs that tells
	// Gnorm how to render and output its schema info.  Each template will be
	// rendered with each schema in turn and written out to the given output
	// path. If no pairs are specified, schemas will not be rendered.
	//
	// The schema path may be a template, in which case the value .Schema may be
	// referenced, containing the name of the current schema being rendered. For
	// example, "schemas/{{.Schema}}/{{.Schema}}.go" = "schemas.gotmpl" would
	// render schemas.gotmpl template with the "public" schema and output to
	// ./schemas/public/public.go
	SchemaPaths map[string]string

	// EnumPaths is a set of "output-path" = "template-path" pairs that tells
	// Gnorm how to render and output its enum info.  Each template will be
	// rendered with each enum in turn and written out to the given output path.
	// If no pairs are specified, enums will not be rendered.
	//
	// The enum path may be a template, in which case the values .Schema and
	// .Enum may be referenced, containing the name of the current schema and
	// Enum being rendered.  For example, "gnorm/{{.Schema}}/enums/{{.Enum}}.go"
	// = "enums.gotmpl" would render the enums.gotmpl template with data from
	// the "public.book_type" enum to ./gnorm/public/enums/users.go.
	EnumPaths map[string]string

	// TypeMap is a mapping of database type names to replacement type names
	// (generally types from your language for deserialization).  Types not in
	// this list will remain in their database form.  In the data sent to your
	// template, this is the Column.Type, and the original type is in
	// Column.OrigType.  Note that because of the way tables in TOML work,
	// TypeMap and NullableTypeMap must be at the end of your configuration
	// file.
	TypeMap map[string]string

	// NullableTypeMap is a mapping of database type names to replacement type
	// names (generally types from your language for deserialization)
	// specifically for database columns that are nullable.  Types not in this
	// list will remain in their database form.  In the data sent to your
	// template, this is the Column.Type, and the original type is in
	// Column.OrigType.   Note that because of the way tables in TOML work,
	// TypeMap and NullableTypeMap must be at the end of your configuration
	// file.
	NullableTypeMap map[string]string

	// Params contains any data you may want to pass to your templates.  This is
	// a good way to make templates reusable with different configuration values
	// for different situations.  The values in this field will be available in
	// the .Params value for all templates.
	Params map[string]interface{}

	// PluginDirs a set of absolute/relative  paths that will be used for
	// plugin lookup.
	PluginDirs []string
}
