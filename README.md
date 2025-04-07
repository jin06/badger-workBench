# Badger Workbench

Badger Workbench is a graphical user interface (GUI) application built with [Fyne](https://fyne.io/) for managing a [BadgerDB](https://dgraph.io/badger) database. It allows users to interact with key-value pairs in the database, including adding, updating, deleting, and viewing keys and their values.

## Features

- **Load Database**: Open a BadgerDB directory to interact with the database.
- **View Keys**: Display all keys stored in the database.
- **Add/Update Key-Value Pairs**: Add new key-value pairs or update existing ones.
- **Set TTL**: Specify a time-to-live (TTL) for keys.
- **Delete Keys**:
  - Delete a selected key.
  - Delete all keys in the database.
- **View TTL**: Display the remaining TTL for keys.
- **Dark Theme**: The application uses a dark theme for better readability.

## Prerequisites

- Go 1.18 or later
- Fyne library (`go get fyne.io/fyne/v2`)
- BadgerDB library (`go get github.com/dgraph-io/badger/v4`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/badger-workbench.git
   cd badger-workbench
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Usage

1. **Load Database**: Click the "Load Badger 🤘" button to select a BadgerDB directory.
2. **View Keys**: Keys will be displayed in the list on the left.
3. **Add/Update Key-Value**:
   - Enter a key in the "Enter Key" field.
   - Enter a value in the "Enter Value" field.
   - Optionally, specify a TTL in seconds in the "Enter TTL (seconds)" field.
   - Click "Submit" to save the key-value pair.
4. **Delete Keys**:
   - To delete a selected key, click "Delete Selected Key 🗑️".
   - To delete all keys, click "Delete All Keys 🗑️".
5. **Refresh Keys**: Click "Refresh Keys 🔁" to reload the key list.

## Project Structure

```
badger-workbench/
├── main.go          # Entry point of the application
├── ui.go            # UI setup and event handling
├── db.go            # Database operations (BadgerDB)
├── README.md        # Project documentation
```

## Screenshots

![Screenshot](https://via.placeholder.com/800x600?text=Badger+Workbench+Screenshot)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fyne](https://fyne.io/) for the GUI framework.
- [BadgerDB](https://dgraph.io/badger) for the key-value database.