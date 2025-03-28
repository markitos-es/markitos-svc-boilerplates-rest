#!/bin/bash
# filepath: /home/markitos/desarrollo/github/public/markitos-svc-boilerplates/bin/clone-caas.sh

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
echo "#  🥷 Clone As A Service (CaaS) 🗡️"
echo "#  Markitos DevSecOps Kulture"
echo "# ============================================="
echo 

# Check if we have the necessary arguments
if [ $# -ne 2 ]; then
    echo "#:[.'.]:> ❌ Error: Se necesitan exactamente 2 argumentos."
    echo "#:[.'.]:> Uso: $0 <nuevo-nombre-servicio> <nombre-entidad>"
    echo "#:[.'.]:> Ejemplo: $0 pepito-svc-mariposas butterfly"
    exit 1
fi

# Assign arguments to variables
NEW_SERVICE_NAME=$1
ENTITY_NAME=$2
ENTITY_NAME_PLURAL="${ENTITY_NAME}s"

# Obtener el nombre del directorio actual (fuente)
CURRENT_DIR=$(basename $(pwd))
SOURCE_DIR="."
TEMP_DIR="/tmp/${NEW_SERVICE_NAME}"
PARENT_DIR=$(dirname $(pwd))
DESTINATION_DIR="${PARENT_DIR}/${NEW_SERVICE_NAME}"

echo
echo "#:[.'.]:> 🚀 Resumen de la operación:"
echo "#:[.'.]:> ----------------------------------------------"
echo "#:[.'.]:> 📂 Directorio original: $CURRENT_DIR (directorio actual)"
echo "#:[.'.]:> 📂 Nuevo directorio: ${DESTINATION_DIR}"
echo "#:[.'.]:> 🔄 Reemplazos que se realizarán:"
echo "#:[.'.]:>   - '$CURRENT_DIR' → '$NEW_SERVICE_NAME'"
echo "#:[.'.]:>   - 'boilerplate' → '$ENTITY_NAME'"
echo "#:[.'.]:>   - 'boilerplates' → '$ENTITY_NAME_PLURAL'"
echo "#:[.'.]:> ----------------------------------------------"
echo "#:[.'.]:> ⚠️  IMPORTANTE: Esta operación NO modificará el directorio original."
echo "#:[.'.]:>    La nueva copia se creará en: ${DESTINATION_DIR}"
echo

read -p "#:[.'.]:> Presiona ENTER para continuar o CTRL+C para cancelar..." CONFIRM

# Clone the directory to temp location first
echo "#:[.'.]:> 🔄 Copiando el directorio actual a ubicación temporal..."
mkdir -p "$TEMP_DIR"
cp -a "$SOURCE_DIR"/* "$TEMP_DIR"
cp -a "$SOURCE_DIR"/.[^.]* "$TEMP_DIR" 2>/dev/null || true

# Remove .git directory if it exists
echo "#:[.'.]:> 🧹 Eliminando el directorio .git por seguridad..."
rm -fr "$TEMP_DIR/.git" 2>/dev/null || true

# Replace content in files
echo "#:[.'.]:> 📝 Reemplazando contenido en todos los archivos..."
find "$TEMP_DIR" -type f -not -path "*/\.*" -exec grep -l "$CURRENT_DIR" {} \; | xargs -I{} sed -i "s/$CURRENT_DIR/$NEW_SERVICE_NAME/g" {} 2>/dev/null || true
find "$TEMP_DIR" -type f -not -path "*/\.*" -exec grep -l "boilerplates" {} \; | xargs -I{} sed -i "s/boilerplates/$ENTITY_NAME_PLURAL/g" {} 2>/dev/null || true
find "$TEMP_DIR" -type f -not -path "*/\.*" -exec grep -l "boilerplate" {} \; | xargs -I{} sed -i "s/boilerplate/$ENTITY_NAME/g" {} 2>/dev/null || true

# Move from temp to final destination
echo "#:[.'.]:> 📦 Moviendo a ubicación final..."
mkdir -p "$(dirname "$DESTINATION_DIR")"
mv "$TEMP_DIR" "$DESTINATION_DIR"

echo "#:[.'.]:> ✅ Proceso completado exitosamente! 🎉"
echo "#:[.'.]:> Tu nuevo servicio está disponible en: $DESTINATION_DIR"
echo "#:[.'.]:> Recuerda crear un nuevo repositorio git si lo necesitas:"
echo "#:[.'.]:> cd $DESTINATION_DIR && git init"
echo