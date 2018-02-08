package tech_curricullum

const getData = `
	SELECT 
		user_id,
		full_name,
		msisdn,
		user_email,
		birth_date,
		create_time,
		update_time 
	FROM (
		SELECT * FROM ws_user ORDER BY full_name DESC LIMIT 1000
	) sub
	ORDER BY full_name ASC limit 10
`
