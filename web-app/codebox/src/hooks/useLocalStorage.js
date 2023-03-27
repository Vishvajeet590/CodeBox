import {useState} from "react";

export function useLocalStorage(key,initialValue) {
    const [storedValue, setStoredValue] = useState(() => {
        if (typeof window === "undefined") {
            return initialValue;
        }
        try {
            const item = window.localStorage.getItem(key);
            if (item === null) {
                let rand =  crypto.randomUUID()
                localStorage.setItem(key, JSON.stringify(rand))
                return rand
            }
            return item ? JSON.parse(item) : initialValue;
        } catch (error) {
            console.log(error);
            return initialValue;
        }
    });
    return [storedValue];
}