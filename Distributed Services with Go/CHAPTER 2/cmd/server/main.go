// Guarantees type-safety;
// Prevents schema-violations;
// Enables fast serialization; and
// Offers backward compatibility;

// Consistent schemas
// My colleagues and I built the infrastructures at my last two companies
// on microservices, and we had a repo called “structs” that housed our
// protobuf and their compiled code, which all our services depended on.
// By doing this, we ensured that we didn’t send multiple, inconsistent
// schemas to prod. Thanks to Go’s type checking, we could update our
// structs dependency, run the tests that touched our data models, and the compiler and tests would tell us whether our code was consistent with
// our schema.

// Versioning for free
