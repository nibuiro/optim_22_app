package model

type Request struct{
	request_id int
	client_id int
	engineer_id list
	name string
	content string
	winner_id int
	finish bool
}