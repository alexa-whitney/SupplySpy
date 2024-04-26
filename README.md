[![Go Report Card](https://goreportcard.com/badge/github.com/alexa-whitney/SupplySpy)](https://goreportcard.com/report/github.com/alexa-whitney/SupplySpy)

# SupplySpy: Streamlined Inventory Management

Welcome to **SupplySpy**, the ultimate solution for managing your office supplies and tech gadgets! Say goodbye to the chaos of overstocking and stockouts. With SupplySpy, efficiency is just a click away!

## Description

**SupplySpy** is a web-based inventory management system built using the Go programming language and the Gin web framework. It's designed to be simple, efficient, and easy to use, allowing users to manage inventory seamlessly. Whether you're tracking staplers or the latest tech gadgets, SupplySpy has you covered.

## Key Features

- **User-Friendly Interface:** Add, remove, and update items with ease.
- **Real-Time Updates:** See changes without needing to reload pages.
- **Persistent Storage:** Your data stays safe and sound in a JSON file.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/SupplySpy.git
   cd SupplySpy
   ```

2. **Set up Go (if not already installed):**
Follow the instructions based on your operating system from the [official GO website](https://go.dev/doc/install).

3. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

Once SupplySpy is up and running, navigate to `http://localhost:8080` in your web browser to start managing your inventory. Here's how you can use the system:

### Viewing Inventory

- **Access the Inventory List:** Simply click on the prompt **To view existing inventory**.

![Screenshot for Inventory List - Main Page](https://i.imgur.com/r79aWu5.png)
![Screenshot for Inventory List - Inventory Page](https://i.imgur.com/cCv4krn.png)

### Adding an Item

- **Navigate to Add Item:** lick on the prompt **To add a new inventory item**.
- **Enter Item Details:** Fill in the ID, Name, Description, and Quantity.
- **Submit:** Click the **Add Item** button to save the new item to your inventory.

![Screenshot for Adding an Item - Main Page](https://i.imgur.com/AP0yLMd.png)
![Screenshot for Adding an Item - Add Item Page](https://i.imgur.com/76LA3NL.png)

### Updating an Item

- **Find the Item:** Each item has an **Edit** button in the inventory list.
- **Modify the Details:** Change the name, description, or quantity directly in the list.
- **Save Changes:** Click the **Save** button that appears after you start editing.

![Screenshot for Updating an Item](https://i.imgur.com/DpQeFYM.png)

### Deleting an Item

- **Delete with Ease:** Next to each item, there's a **Delete** button. Just click it, and the item is removed from your inventory.

![Screenshot 1 for Deleting an Item](https://i.imgur.com/isluJNs.png)
![Screenshot 2 for Deleting an Item](https://i.imgur.com/7ynW3mu.png)


## Running Tests

To run the table-driven tests, run the following command

```bash
  go test
```
![Screenshot for passing table tests](https://i.imgur.com/Z4xu2kc.png)

To run the benchmark test, run the following command

```bash
  go test -bench=.
```
![Screenshot for passing benchmark tests](https://i.imgur.com/Ulc1WKg.png)

## Contributing

Feel free to fork the repository and submit pull requests. You can also open issues if you encounter bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Final Notes

Enjoy using SupplySpy for all your inventory management needs! Enhance productivity and reduce costs effectively. For more information or assistance, contact alexawhitney87@gmail.com.
