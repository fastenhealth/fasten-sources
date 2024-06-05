export function uuidV4(): string {
    // https://stackoverflow.com/questions/55853005/error-ts2365-operator-cannot-be-applied-to-types-number-and-1000
    return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, c =>
        (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
}
