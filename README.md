# MiRoutine

Aplicación web para gestionar rutinas de gimnasio, con backend en Go y frontend en Vue.js.

## Tecnologías principales

- **Frontend:** Vue.js, Tailwind CSS, Vite
- **Backend:** Go (Golang), GORM, PostgreSQL

## Estructura del proyecto

```
├── gym-frontend/   # Frontend (Vue.js)
├── gym-backend/    # Backend (Go)
└── pasos.txt       # Notas o pasos personalizados
```

## Instalación y uso

### Backend (Go)

1. Entra a la carpeta del backend:
   ```bash
   cd gym-backend
   ```
2. Configura tu base de datos PostgreSQL y ajusta el DSN en `database/db.go` si es necesario.
3. Instala las dependencias y ejecuta:
   ```bash
   go mod tidy
   go run main.go
   ```

### Frontend (Vue.js)

1. Entra a la carpeta del frontend:
   ```bash
   cd gym-frontend
   ```
2. Instala las dependencias:
   ```bash
   npm install
   ```
3. Inicia la app:
   ```bash
   npm run dev
   ```

## Funcionalidades principales

- Registro e inicio de sesión de usuarios
- Visualización de rutinas personalizadas
- Panel protegido para usuarios autenticados

## Notas

- Asegúrate de tener PostgreSQL corriendo y configurado correctamente.
- El frontend y backend corren por separado y se comunican vía API REST.

---

¡Contribuciones y sugerencias son bienvenidas!
