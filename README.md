# Echo_GO_Framework

![Go](https://img.shields.io/badge/Go-1.23+-00add8?style=for-the-badge&logo=go&logoColor=white)&nbsp;![Echo](https://img.shields.io/badge/Echo-v4.13.3-00add8?style=for-the-badge&logo=go&logoColor=white)&nbsp;![HTML Templates](https://img.shields.io/badge/HTML-Templates-e34f26?style=for-the-badge&logo=html5&logoColor=white)&nbsp;![Estado](https://img.shields.io/badge/Estado-En%20progreso-22c55e?style=for-the-badge)

> **Echo_GO_Framework** es una aplicación web en **Go** construida con el framework **Echo v4**, que implementa enrutamiento dinámico, renderizado de plantillas HTML nativas, servicio de archivos estáticos y middleware de logging personalizado.

---

## 📋 Descripción

El proyecto sirve como base de aprendizaje para construir aplicaciones web con Go y Echo. Incluye:

- **Enrutamiento dinámico** con parámetros de ruta (`/:page`, `/:nombre`).
- **Renderizado de plantillas** HTML nativas con `html/template` y `ParseGlob`.
- **Archivos estáticos** servidos desde la carpeta `/static`.
- **Middleware de logging** personalizado con formato de método, URI, status y latencia.
- **Manejo de 404** con plantilla propia cuando la página no existe.

---

## 🛣️ Endpoints disponibles

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/` | Página principal con plantilla `index.html` |
| `GET` | `/:page` | Renderiza cualquier plantilla HTML dinámicamente |
| `GET` | `/saludo` | Devuelve un saludo estático en texto plano |
| `GET` | `/saludo/:nombre` | Devuelve un saludo personalizado con el nombre |

---

## 🏗️ Estructura del Proyecto

```txt
Echo_GO_Framework/
├── routes/
│   └── routes.go          # Definición de rutas y renderer de plantillas
├── static/                # Archivos estáticos (CSS, JS, imágenes)
├── templates/             # Plantillas HTML (index.html, 404.html, ...)
├── main.go                # Punto de entrada: instancia Echo y arranca en :8080
├── go.mod                 # Módulo Go y dependencias
├── go.sum                 # Checksums de dependencias
└── README.md              # Documentación del proyecto
```

---

## ⚙️ Instalación y Ejecución

Clonar el repositorio:
```bash
git clone https://github.com/sorgazb/Echo_GO_Framework.git
cd Echo_GO_Framework
```

Descargar dependencias:
```bash
go mod tidy
```

Levantar el servidor en desarrollo:
```bash
go run main.go
```

El servidor estará disponible en:
```txt
http://localhost:8080
```

---

## 🤝 Contribución

Haz fork del repositorio.

Crea una rama de trabajo:
```bash
git checkout -b feature/mi-nueva-funcionalidad
```

Realiza tus cambios y haz commit.

Abre un Pull Request describiendo tus mejoras.

---

<p align="center">Aprendizaje de Echo Framework &nbsp;·&nbsp; Sergio Orgaz Bravo</p>
