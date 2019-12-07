#!/bin/bash
set -m

eval "$@" &

# Wait 60 seconds for SQL Server to start up by ensuring that 
# calling SQLCMD does not return an error code, which will ensure that sqlcmd is accessible
# and that system and user databases return "0" which means all databases are in an "online" state
# https://docs.microsoft.com/en-us/sql/relational-databases/system-catalog-views/sys-databases-transact-sql?view=sql-server-2017 

DBSTATUS=1
ERRCODE=1
declare -i i=0

while [[ "${DBSTATUS}${ERRCODE}" != "00" ]] && [[ $i < 60 ]]; do
	i+=1
	DBSTATUS=$(/opt/mssql-tools/bin/sqlcmd -h -1 -t 1 -U sa -P $SA_PASSWORD -Q "SET NOCOUNT ON; Select SUM(state) from sys.databases" | grep -o "[0-9]*")
	ERRCODE=$?
	sleep 1
	echo "Waiting for sql server to come online $i"
done

if [[ "${DBSTATUS}${ERRCODE}" != "00" ]]; then 
	echo "SQL Server took more than 60 seconds to start up or one or more databases are not in an ONLINE state. DBSTATUS: ${DBSTATUS} ERRCODE: ${ERRCODE}"
	exit 1
fi

# Run the setup script to create the DB and the schema in the DB

DATA_DIR=/usr/config/setup_data
DBS="${DATA_DIR}/*"

for d in $DBS; do
	db=$(basename $d)
	/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $SA_PASSWORD -Q "CREATE DATABASE ${db}"
	for f in ${d}/*.sql; do
		echo "executing script $f in db $db ..."
		[ -e "$f" ] || continue
		/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $SA_PASSWORD -d $db -i "$f"
	done
done
fg 1