# Distributed Log Package

## Overview

This package provides an implementation of a distributed log system that stores data in a series of organized records, files, and segments.

### Key Components

- **Record**: A unit of data stored within the log.
- **Store**: A file where record data is saved.
- **Index**: A file that stores references to records for efficient lookup.
- **Segment**: Combines a store and index, managing a collection of records.
- **Log**: Aggregates all segments, serving as the complete storage unit for records.

### Functionality

- **Reading a Record by Offset**
  - Retrieves the record’s position from the memory-mapped index file.
  - Reads the record directly from the store file at the retrieved position.

This structure enables efficient data storage and retrieval for distributed systems.
