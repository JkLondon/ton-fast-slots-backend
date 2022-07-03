package repository

const (
	queryCreateUser = `INSERT
	INTO public.app_user
		(tg_id, wallet_address, wallet_seed)
	VALUES
		($1, $2, $3);`
	queryGetSelfInfo = `SELECT au.tg_id, au.wallet_address FROM app_user au WHERE au.tg_id = $1;`
)
