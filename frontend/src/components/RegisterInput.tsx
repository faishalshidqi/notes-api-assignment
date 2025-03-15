import useStringInput from '../hooks/hooks.ts'

export default function RegisterInput({onRegister}: {onRegister: ({fullname, username, password}: {fullname: string, username: string, password: string}) => void}) {
    const [name, setName] = useStringInput()
    const [username, setUsername] = useStringInput()
    const [password, setPassword] = useStringInput()
    const [confirmPassword, setConfirmPassword] = useStringInput()
    const onSubmitHandler = (event: {preventDefault: () => void}) => {
        event.preventDefault()
        if (password !== confirmPassword) {
            alert('Passwords do not match')
            return
        }
        onRegister({fullname: name as string, username: username as string, password: password as string})
    }
    return (
        <div className='input-register'>
            <form onSubmit={onSubmitHandler}>
                <label htmlFor='name'>Name</label>
                <input type='text' name='name' placeholder='Name' value={name as string} onChange={setName as (event: {target: {value: string}}) => void} />
                <label htmlFor='username'>Username</label>
                <input type='username' name='username' placeholder='username' value={username as string} onChange={setUsername as (event: {target: {value: string}}) => void} />
                <label htmlFor='password'>Password</label>
                <input type='password' name='password' placeholder='Password' value={password as string} onChange={setPassword as (event: {target: {value: string}}) => void} />
                <label htmlFor='confirmPassword'>Confirm Password</label>
                <input type='password' name='password' placeholder='Confirm Password' value={confirmPassword as string} onChange={setConfirmPassword as  (event: {target: {value: string}}) => void} />
                <button type='submit'>Submit</button>
            </form>
        </div>
    )
}