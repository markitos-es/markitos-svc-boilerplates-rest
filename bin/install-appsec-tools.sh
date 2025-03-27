#!/bin/bash

echo 
echo "#:[.'.]:>=============================================="
echo "#:[.'.]:>  __  __  ____  _  __"
echo "#:[.'.]:> |  \\/  |  _ \\| |/ /"
echo "#:[.'.]:> | \\  / | | | | ' / "
echo "#:[.'.]:> | |\\/| | | | |  <  "
echo "#:[.'.]:> | |  | | |_| | . \\ "
echo "#:[.'.]:> |_|  |_|____/|_|\\_\\"
echo "#:[.'.]:>                                   "
echo "#:[.'.]:>  Creador: Marco Antonio - markitos      "
echo "#:[.'.]:>=============================================="
echo "#:[.'.]:>= 🥷 (mArKit0sDevSecOpsKit) 🗡️"
echo "#:[.'.]:>= Markitos DevSecOps Kulture"
echo "#:[.'.]:>=============================================="
echo 

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"
set -euo pipefail
IFS=$'\n\t'

#:[.'.]:> Definir funciones de logging
function log_info() {
    echo -e "\033[1;34m[INFO]\033[0m $*"
}

function log_error() {
    echo -e "\033[1;31m[ERROR]\033[0m $*" >&2
}

function log_success() {
    echo -e "\033[1;32m[SUCCESS]\033[0m $*"
}

#:[.'.]:> Detectar si SNYK_TOKEN ya está configurado
if [[ -n "${SNYK_TOKEN:-}" ]]; then
    DETECTED_SNYK_TOKEN="${SNYK_TOKEN:0:8}****"
    echo -e "\033[1;36m🔑 Hemos detectado un SNYK_TOKEN configurado: $DETECTED_SNYK_TOKEN\033[0m"
    echo -e "\033[1;36mSi deseas usar otro, introdúcelo a continuación. Si no, presiona ENTER para continuar con el detectado.\033[0m"
else
    echo -e "\033[1;36m🔑 No se detectó un SNYK_TOKEN configurado. Introduce uno a continuación o presiona ENTER para usar 'replace_me'.\033[0m"
fi
if [[ -z "${ASK_FOR_SNYK_TOKEN_BYPASS:-}" || "${ASK_FOR_SNYK_TOKEN_BYPASS}" != "true" ]]; then
    read -p "Introduce tu SNYK_TOKEN: " INPUT_SNYK_TOKEN
    SNYK_TOKEN=${INPUT_SNYK_TOKEN:-${SNYK_TOKEN:-replace_me}}
else
    echo -e "\033[1;36m🔑 ASK_FOR_SNYK_TOKEN_BYPASS está habilitado. Usando el SNYK_TOKEN existente o 'replace_me'.\033[0m"
    SNYK_TOKEN=${SNYK_TOKEN:-replace_me}
fi

#:[.'.]:> Mostrar lo que hará el script
echo -e "\033[1;36m🛠️ Este script instalará las siguientes herramientas en ~/.local/bin:\033[0m"
echo -e "  - \033[1;33mSnyk CLI\033[0m (https://snyk.io)"
echo -e "  - \033[1;33mGitleaks\033[0m (https://github.com/gitleaks/gitleaks)"
echo -e "  - \033[1;33mPre-commit hook\033[0m para seguridad en Git."
echo
echo -e "\033[1;36m📋 Resumen de acciones:\033[0m"
echo -e "  1. Descargar los binarios."
echo -e "  2. Moverlos a ~/.local/bin."
echo -e "  3. Hacerlos ejecutables."
echo -e "  4. Instalar el pre-commit hook."
echo -e "  5. Actualizar el PATH y configurar SNYK_TOKEN."
echo -e "  6. Verificar las versiones instaladas."
echo
echo -e "\033[1;33m⚠️ Presiona CTRL+C para cancelar. Si no, el script continuará automáticamente en 10 segundos...\033[0m"
sleep 10

#:[.'.]:> Crear directorio ~/.local/bin si no existe
mkdir -p ~/.local/bin

#:[.'.]:> Instalar Snyk CLI
log_info "Descargando e instalando Snyk CLI..."
cd /tmp
curl -s https://static.snyk.io/cli/latest/snyk-linux -o snyk
chmod +x ./snyk
mv ./snyk ~/.local/bin/snyk
chmod u+x ~/.local/bin/snyk

#:[.'.]:> Instalar Gitleaks
log_info "Descargando e instalando Gitleaks..."
cd /tmp
wget -q https://github.com/gitleaks/gitleaks/releases/download/v8.24.0/gitleaks_8.24.0_linux_x64.tar.gz
tar xfz gitleaks_8.24.0_linux_x64.tar.gz
mv ./gitleaks ~/.local/bin/gitleaks
chmod u+x ~/.local/bin/gitleaks

#:[.'.]:> Instalar pre-commit hook
cd "$SCRIPT_DIR/../"
log_info "Instalando pre-commit hook..."
PRE_COMMIT_SOURCE="$(pwd)/etc/pre-commit"
PRE_COMMIT_DESTINATION=".git/hooks/pre-commit"

if [ -f "$PRE_COMMIT_SOURCE" ]; then
    cp "$PRE_COMMIT_SOURCE" "$PRE_COMMIT_DESTINATION"
    chmod u+x "$PRE_COMMIT_DESTINATION"
    log_success "Pre-commit hook instalado correctamente en $PRE_COMMIT_DESTINATION"
else
    log_error "No se encontró el archivo pre-commit en $PRE_COMMIT_SOURCE"
fi

#:[.'.]:> Actualizar PATH y configurar SNYK_TOKEN
if ! grep -Fxq 'export PATH=${PATH}:${HOME}/.local/bin' ~/.bashrc; then
    log_info "Actualizando PATH en ~/.bashrc..."
    echo 'export PATH=${PATH}:${HOME}/.local/bin' >> ~/.bashrc
    log_success "PATH actualizado correctamente."
else
    log_info "El PATH ya incluye ~/.local/bin. No se realizaron cambios."
fi

if ! grep -Fxq "export SNYK_TOKEN=${SNYK_TOKEN}" ~/.bashrc; then
    log_info "Configurando SNYK_TOKEN en ~/.bashrc..."
    echo "export SNYK_TOKEN=${SNYK_TOKEN}" >> ~/.bashrc
    log_success "SNYK_TOKEN configurado correctamente."
else
    log_info "El SNYK_TOKEN ya está configurado en ~/.bashrc. No se realizaron cambios."
fi

source ~/.bashrc

#:[.'.]:> Verificar las herramientas instaladas
log_info "Verificando las herramientas instaladas..."
SNYK_VERSION=$(~/.local/bin/snyk --version 2>/dev/null || echo "No instalado")
GITLEAKS_VERSION=$(~/.local/bin/gitleaks version 2>/dev/null || echo "No instalado")

#:[.'.]:> Mostrar informe final
echo
echo -e "\033[1;36m📋 Informe final:\033[0m"
if [[ "$SNYK_VERSION" != "No instalado" ]]; then
    log_success "Snyk CLI instalado correctamente. Versión: $SNYK_VERSION"
else
    log_error "Snyk CLI no se instaló correctamente."
fi

if [[ "$GITLEAKS_VERSION" != "No instalado" ]]; then
    log_success "Gitleaks instalado correctamente. Versión: $GITLEAKS_VERSION"
else
    log_error "Gitleaks no se instaló correctamente."
fi

if [ -f "$PRE_COMMIT_DESTINATION" ]; then
    log_success "Pre-commit hook instalado correctamente."
else
    log_error "El pre-commit hook no se instaló correctamente."
fi

#:[.'.]:> Limpiar archivos temporales
log_info "Limpiando archivos temporales..."
rm -f /tmp/gitleaks_8.24.0_linux_x64.tar.gz
rm -f /tmp/README.md /tmp/LICENSE

echo
log_success "🎉 Instalación completada. ¡Todo listo para usar!"