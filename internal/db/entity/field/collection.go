package field

var (
	CollectionProblem = "tbl_collection.id IN (SELECT collection_id FROM tbl_collection_problem WHERE problem_id In(?) GROUP BY collection_id HAVING COUNT(DISTINCT problem_id) =?)"
)
