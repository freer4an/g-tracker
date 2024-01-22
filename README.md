# G-Tracker

## Overview

G-Tracker is a pet project written in Go, HTML, CSS, and JS that provides a minimalistic web interface for tracking information about music artists and bands. The project utilizes the GroupieTrackers API (https://groupietrackers.herokuapp.com/api) to fetch data about music artists and bands. The application supports various features such as filters, sorts, and search functionality to enhance user experience.

## Features

- **API Integration:** G-Tracker parses the GroupieTrackers API to retrieve information about music artists and bands.

- **Filters/Sorts/Search:** The application allows users to filter, sort, and search for specific artists or bands, providing a tailored experience.

- **Singleton Pattern:** The project adopts the singleton pattern to ensure a single instance of the API parser, promoting efficiency and resource management.

- **Ticker for API Updates:** G-Tracker utilizes a ticker to periodically check if the API has been updated. This ensures that the application remains up-to-date with the latest information.

- **Standard Go Libraries:** The entire project is built using only standard Go libraries, minimizing dependencies and simplifying the development and deployment process.

- **Custom Router (Experimental):** A custom router has been implemented, to handle routes and support dynamic route parameters. Please note that this router is experimental and not claimed to be best practice. It was implemented for experimentation purposes.

- **Simple Custom Logger:** G-Tracker includes a straightforward custom logger for logging events and messages, aiding in debugging and monitoring.


