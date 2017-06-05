package postgresql

const (
	modelTable = "model_table"
)

/*
 * Database select queries used in the package to generate
 * prepared statements for transactions.
 */
const (
	selectModelBase = `
	SELECT t.id 
	FROM ` + modelTable + " AS t "
	selectModelByID = selectModelBase + "WHERE t.id = $1"
)

/*
 * Database insert queries for model
 */
const (
	insertModel = `
	INSERT INTO ` + modelTable + `(id)
	VALUES (:id)
	RETURNING id
	`
)

/*
 * Database update queries
 */
const (
	updateModelBase = `
	UPDATE ` + modelTable + ` AS t
	SET `
	disableModelBase = updateModelBase + "enabled = false "
	disableModelByID = disableModelBase + "WHERE id=:id"
)
