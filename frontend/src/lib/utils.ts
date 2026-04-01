export async function snoop<T>(promise: Promise<T>): Promise<[T | null, string | null]> {
    try {
        const result = await promise;
        return [result, null];
    } catch (err: any) {
        return [null, String(err)];
    }
}