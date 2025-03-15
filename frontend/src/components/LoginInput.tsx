import useStringInput from '../hooks/hooks.ts'

export default function LoginInput({onLogin}: {onLogin: ({username, password}: {username: string, password: string}) => void}) {
    const [username, setUsername] = useStringInput()
    const [password, setPassword] = useStringInput()
    const onSubmitHandler = (event: { preventDefault: () => void }) => {
        event.preventDefault()
        onLogin({username: username as string, password: password as string})
    }
    return (
        <div className='input-login'>
            <form onSubmit={onSubmitHandler}>
                <label htmlFor='username'>Username</label>
                <input type='username' name='Username' placeholder='Username' value={username as string} onChange={setUsername as (event: {target: {value: string}}) => void} />
                <label htmlFor='password'>Password</label>
                <input type='password' name='password' placeholder='Password' value={password as string} onChange={setPassword as (event: {target: {value: string}}) => void} />
                <button type='submit'>Submit</button>
            </form>
        </div>
    )
}