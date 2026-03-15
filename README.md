# ⚡ InstantCheck

**InstantCheck** es un monitor de estado HTTP ultra rápido escrito en Go. Utiliza concurrencia nativa (Goroutines y Channels) para verificar la disponibilidad de múltiples servidores de forma simultánea, ahorrando tiempo y recursos.



## ✨ Características

- **Verificación Concurrente:** No importa si tienes 5 o 500 URLs, se chequean todas al mismo tiempo.
- **Lectura desde Archivo:** Gestiona tus URLs fácilmente en un archivo de texto simple.
- **Ligero:** Un binario único sin dependencias externas.
- **Informativo:** Reporta códigos de estado HTTP y detecta fallos de conexión.

## 🚀 Instalación

Asegúrate de tener [Go](https://go.dev/) instalado en tu sistema, luego clona este repositorio y compila:

```bash
git clone [https://github.com/tu-usuario/instantcheck](https://github.com/tu-usuario/instantcheck)
cd instantcheck
go build -o instantcheck
```

## Pictures 📷
<img width="630" height="539" alt="image" src="https://github.com/user-attachments/assets/1bd22a1c-5d02-4dd6-9a46-c847a7813bf0" />


## Compilar 👌

**Ejecutar en la terminal**
```bash
 GOOS=linux go build main.go instantcheck
```
