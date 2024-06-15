package utils

import (
	"fmt"
	"strings"
)

func GeneratePostgresAlterColumnIfExistStatement(tableName string, columns []string, alterExpression string) string {
	var script strings.Builder

	script.WriteString(`
    DO $$
    DECLARE
        current_column text;
    BEGIN`)

	for _, column := range columns {
		script.WriteString(fmt.Sprintf(`
        current_column := '%s';
        IF EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = '%s' 
            AND column_name = current_column
        ) THEN
            EXECUTE format('ALTER TABLE %%I ALTER COLUMN %%I %s', '%s', current_column);
        ELSE
            RAISE NOTICE 'Column %% does not exist or alteration condition not met.', current_column;
        END IF;`, column, tableName, alterExpression, tableName))
	}

	script.WriteString(`
    END$$;
    `)

	return script.String()
}
