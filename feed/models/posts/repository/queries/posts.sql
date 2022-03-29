
-- name: GetPostingByID :one
SELECT * FROM posting 
WHERE id = $1;

-- name: GetAllPostings :many
SELECT * FROM posting;

-- name: GetCompanyJobs :many
SELECT * FROM posting 
WHERE company_name = $1; 

-- name: GetJobs :many
SELECT * FROM posting 
WHERE company_name = $1 
AND title = $2
AND salary >= $3; 


-- name: DeletePosting :one
DELETE FROM posting
WHERE id = $1
RETURNING *;


-- name: CreateNewPosting :one
INSERT INTO posting(title, company_name, company_description, 
posting_description, salary)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
    
)
RETURNING *;


-- name: GetApplicants :many
SELECT * FROM users INNER JOIN applicant_info
ON users.id = applicant_info.applicant_id
WHERE company_name = $1;