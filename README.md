# Hot Beverages Recipe Finder

This project is a simple web application for searching and displaying recipes for popular hot beverages like Espresso, Americano, and Latte. It includes a web interface, a backend in Go, and tests to ensure functionality.

## Features

- **Search Recipes**: Users can search for a recipe by entering the name of a beverage.
- **Dynamic Rendering**: Recipes are dynamically rendered using templates.
- **Static File Support**: The app serves static files such as CSS for styling.
- **Error Handling**: Graceful error handling for unsupported or non-existent recipes.
- **Comprehensive Tests**: Includes test cases for API handlers and static file serving.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/) (version 1.19 or later)
- A web browser to access the application.

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/your-username/Recipe.git
    cd Recipie
    ```

2. Create a directory structure:
    ```bash
    mkdir -p static/templates
    ```

3. Place the following files in their respective directories:
    - `style.css` in `static/`
    - `index.html` in `templates/`

4. Build the application:
    ```bash
    go build
    ```

## Usage

1. Start the server:
    ```bash
    go run main.go
    ```

2. Open your browser and visit:
    ```
    http://localhost:2000
    ```

3. Use the search bar to enter a beverage name (e.g., "Espresso").
 

## Project Structure

```plaintext
.
├── static/
│   └── style.css         # Styles for the application
├── templates/
│   └── index.html        # HTML template for rendering the web page
├── main.go               # Application's main code
├── main_test.go          # Unit tests for the application
└── README.md             # Documentation (this file)

`````

## Testing

Run the test suite to verify application functionality:

```bash
go test -v

`````
## Test Cases

The project includes comprehensive test cases to ensure proper functionality:

### Search Functionality
- Tests various queries to ensure the search handler responds correctly.
- Example:
    - Query: `espresso`
    - Expected Response: `Espresso`
- Handles edge cases for unknown recipes like `matcha` with a fallback response.

### Static File Serving
- Verifies that the application serves static files correctly.
- Tests the following scenarios:
    - Serving existing files like `example.html`.
    - Correct handling of non-existent files with a `404` response.

### Error Handling
- Ensures that the application gracefully handles invalid routes or missing resources.

### MIME Type Verification
- Confirms that files are served with the correct MIME type.
- Example:
    - File: `example.html`
    - Expected MIME: `text/html; charset=utf-8`

## Example Test Outputs

### Search Functionality

For the query `espresso`:
- **Expected Status**: `200 OK`
- **Expected Body**: `Espresso`

For the query `unknown`:
- **Expected Status**: `200 OK`
- **Expected Body**: `Recipe not found`

### Static File Handling

- **Existing File**: `/static/example.html`
    - **Expected Status**: `200 OK`
    - **Expected Body**: Contains `<html>` content.
- **Non-Existent File**: `/static/non-existent-file.txt`
    - **Expected Status**: `404 Not Found`

### MIME Type Verification

- **File**: `/static/example.html`
    - **Expected MIME Type**: `text/html; charset=utf-8`
- **File**: `/static/example.json`
    - **Expected MIME Type**: `application/json`
- **File**: `/static/example.png`
    - **Expected MIME Type**: `image/png`

## Running Tests

To run specific test cases or view detailed outputs:

```bash
go test -run <TestName> -v

`````````
## Test Coverage

This application covers:

- API functionality
- Static file serving
- Error handling
- MIME type validation








    
