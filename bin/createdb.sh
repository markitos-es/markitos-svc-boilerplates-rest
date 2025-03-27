echo 
echo "# ============================================="
echo "#  __  __  ____  _  __"
echo "# |  \\/  |  _ \\| |/ /"
echo "# | \\  / | | | | ' / "
echo "# | |\\/| | | | |  <  "
echo "# | |  | | |_| | . \\ "
echo "# |_|  |_|____/|_|\\_\\"
echo "#                                   "
echo "#  Creador: Marco Antonio - markitos      "
echo "# ============================================="
echo "#  🥷 (mArKit0sDevSecOpsKit) 🗡️"
echo "#  Markitos DevSecOps Kulture"
echo "# ============================================="
echo 
# go to root of project
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."
set -euo pipefail
IFS=$'\n\t'

# Funciones básicas
function log_info() {
    echo "[INFO] $*"
}

function log_error() {
    echo "[ERROR] $*" >&2
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Configuración de variables de entorno
#:[.'.]:>-------------------------------------
source "$SCRIPT_DIR/environment.sh"
setup_environment
show_config "full"
#:[.'.]:>-------------------------------------


# Extraer datos de conexión de DATABASE_DSN
DB_NAME=$(echo $DATABASE_DSN | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="dbname"){print $(i+1)}}}')
DB_USER=$(echo $DATABASE_DSN | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="user"){print $(i+1)}}}')
DB_PASS=$(echo $DATABASE_DSN | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="password"){print $(i+1)}}}')

# Verificar si la base de datos existe
function database_exists() {
    docker exec markitos-common-postgres psql -U admin -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" | grep -q 1
}

# Verificar si el usuario existe
function user_exists() {
    docker exec markitos-common-postgres psql -U admin -d postgres -tAc "SELECT 1 FROM pg_roles WHERE rolname='$DB_USER'" | grep -q 1
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Tu lógica aquí
#:[.'.]:>-------------------------------------
log_info "Argumentos proporcionados: $*"
log_info "Creando base de datos $DB_NAME"

if database_exists "$DB_NAME"; then
    log_info "La base de datos $DB_NAME ya existe"
else
    docker exec markitos-common-postgres createdb --username=admin --owner=admin "$DB_NAME"
    log_info "Base de datos $DB_NAME creada"
fi

if user_exists "$DB_USER"; then
    log_info "El usuario $DB_USER ya existe"
else
    docker exec markitos-common-postgres psql -U admin -d postgres -c "CREATE USER \"$DB_USER\" WITH PASSWORD '$DB_PASS';"
    log_info "Usuario $DB_USER creado"
fi

docker exec markitos-common-postgres psql -U admin -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE \"$DB_NAME\" TO \"$DB_USER\";"
docker exec markitos-common-postgres psql -U admin -d postgres -c "GRANT ALL PRIVILEGES ON SCHEMA public TO \"$DB_USER\";"
log_info "Privilegios otorgados al usuario $DB_USER en la base de datos $DB_NAME"
#:[.'.]:>-------------------------------------