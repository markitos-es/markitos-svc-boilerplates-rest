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
cd "$SCRIPT_DIR/../"

set -euo pipefail
IFS=$'\n\t'

SCRIPT_NAME=$(basename "$0")
LOG_FILE="/tmp/${SCRIPT_NAME%.sh}.log"

# Funciones básicas
function log_info() {
    echo "[INFO] $*" | tee -a "$LOG_FILE"
}

function log_error() {
    echo "[ERROR] $*" >&2 | tee -a "$LOG_FILE"
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Configuración de variables de entorno
#:[.'.]:>-------------------------------------
source "$SCRIPT_DIR/environment.sh"
setup_environment
show_config "full"
#:[.'.]:>-------------------------------------

go clean -testcache && go test -v ./... 
#:[.'.]:>-------------------------------------