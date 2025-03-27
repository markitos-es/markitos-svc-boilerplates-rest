#!/bin/bash

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
echo "#  🥷 (Cultura DevSecOps) 🗡️"
echo "#  Markitos DevSecOps Kulture. "
echo "# ============================================="
echo 

# Ir al directorio raíz del proyecto
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
DB_NAME=$(echo "$DATABASE_DSN" | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="dbname"){print $(i+1)}}}')
DB_USER=$(echo "$DATABASE_DSN" | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="user"){print $(i+1)}}}')

# Eliminar base de datos y usuario
log_info "Eliminando base de datos y usuario asociados"

docker exec markitos-common-postgres psql -U admin -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='$DB_NAME';" || true
docker exec markitos-common-postgres psql -U admin -d postgres -c "DROP DATABASE IF EXISTS \"$DB_NAME\";" || log_error "Error al eliminar la base de datos $DB_NAME"
log_info "Base de datos $DB_NAME eliminada (si existía)"

docker exec markitos-common-postgres psql -U admin -d postgres -c "DROP USER IF EXISTS \"$DB_USER\";" || log_error "Error al eliminar el usuario $DB_USER"
log_info "Usuario $DB_USER eliminado (si existía)"

log_info "Proceso de eliminación completado"