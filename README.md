# twortle
A word puzzle site and helper.

**twortle** is a web-based application designed to help users solve word puzzles and play word games. It provides tools for searching word patterns, playing games, and drawing puzzle patterns, all powered by a Go backend and a SQLite database.

## Prerequisites

To build and run this application locally, you will need:

*   **Go 1.25** or higher.
*   **GCC and musl-dev** (or equivalent C build tools) for CGO support, as the project uses `go-sqlite3`.
*   **Podman** or **Docker** (optional, if you prefer running via containers).

    ## Quick Start (Container)

    If you just want to run the application using the pre-built image from Quay.io:

    ```bash
    podman run -p 3000:3000 quay.io/cmbsolver/twortle
    ```
    *(Or use `docker` if preferred)*. The application will be accessible at `http://localhost:3000`.

    ## Building the Application

    ### Using the Build Script (Recommended)
    The project includes a `build.sh` script that automates building a container image and starting the application. It currently uses `podman`.
    ```bash
    ./build.sh
    ```
    The application will be accessible at `http://localhost:3000`.

    ### Manual Go Build
    If you prefer to run the binary directly on your host machine:

    1.  **Install dependencies:**
        ```bash
        go mod download
        ```
    2.  **Build the binary:**
        Ensure `CGO_ENABLED=1` is set for SQLite support.
        ```bash
        CGO_ENABLED=1 go build -o twortle main.go
        ```
    3.  **Run the application:**
        ```bash
        ./twortle
        ```

    ## Development Stack
    *   **Backend:** Go (using [Fiber](https://gofiber.io/))
    *   **Database:** SQLite via [GORM](https://gorm.io/)
    *   **Frontend:** HTML templates with Bootstrap 5
