import React, {useEffect, useState} from "react";

export const useHistoryLocalStorage = (keyName, defaultValue) => {
    const [storedValue, setStoredValue] = React.useState(() => {
        try {
            const value = window.localStorage.getItem(keyName);

            if (value) {
                return JSON.parse(value);
            } else {
                window.localStorage.setItem(keyName, JSON.stringify(defaultValue));
                return defaultValue;
            }
        } catch (err) {
            return defaultValue;
        }
    });

    const setValue = newValue => {
        try {
            window.localStorage.setItem(keyName, JSON.stringify(newValue));
        } catch (err) {}
        setStoredValue(newValue);
    };

    return [storedValue, setValue];
};