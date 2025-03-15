import {useState} from 'react'

export default function useStringInput(defaultValue: string = '') {
    const [value, setValue] = useState<string>(defaultValue)
    const onValueChange = (event: { target: { value: string } }) => {
        setValue(event.target.value)
    }
    return [value, onValueChange]
}