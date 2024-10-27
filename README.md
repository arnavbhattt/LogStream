# Distributed Log Package Implementation

- Terms

  - Record: data stored in log
  - Store: file records entries are stored in
  - Index: file index entries are stored in
  - Segment: ties store and index together
  - Log: ties all segments together

- Reading a record given offset
  - Get entry from index file (memory-mapped) for record
  - Read record at position in store file
