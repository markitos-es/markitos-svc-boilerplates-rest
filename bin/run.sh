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

# USO:
#:[.'.]:> En desarrollo:
#:[.'.]:>   ./bin/run.sh                           → Usa valores por defecto
#:[.'.]:>   DATABASE_DSN=... HTTP_SERVER_ADDRESS=... ./bin/run.sh → Personalizado
#:[.'.]:> 
#:[.'.]:> Con make:
#:[.'.]:>   make run                               → Usa valores por defecto
#:[.'.]:>   DATABASE_DSN=... make run              → Personalizado

# go to root of project
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"

set -euo pipefail
IFS=$'\n\t'

#:[.'.]:>-------------------------------------
#:[.'.]:> Configuración de variables de entorno
#:[.'.]:>-------------------------------------
source "$SCRIPT_DIR/environment.sh"
setup_environment
show_config "full"
#:[.'.]:>-------------------------------------

#:[.'.]:> Muestra la configuración que vamos a usar
echo "#:[.'.]:> 🚀 Iniciando con configuración:"
echo "#:[.'.]:> 📊 DATABASE_DSN=$DATABASE_DSN"
echo "#:[.'.]:> 🌐 HTTP_SERVER_ADDRESS=$HTTP_SERVER_ADDRESS"
echo "#:[.'.]:>-------------------------------------"

#:[.'.]:> Ejecuta la aplicación con la configuración establecida
go run cmd/main.go
#:[.'.]:>-------------------------------------