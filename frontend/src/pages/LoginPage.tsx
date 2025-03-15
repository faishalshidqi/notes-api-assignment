import {login} from '../utils/data.ts'
import {Link} from 'react-router-dom'
import LoginInput from '../components/LoginInput'
import LocaleContext from '../contexts/LocaleContext.ts'
import {useContext} from 'react'

export default function LoginPage({onSuccessLogin}: {onSuccessLogin: ({accessToken}: {accessToken: string}) => void}) {
    const localeContext = useContext(LocaleContext)
    async function onLoginHandler({username, password}: {username: string, password: string}) {
        const {error, data} = await login({username, password})
        if (!error) onSuccessLogin(data)
    }
    function returnComponent({heading, linkQuestion, linkMessage}: {heading: string, linkQuestion: string, linkMessage: string}) {
        return (
            <section className={'login-page'}>
                <h2>{heading}</h2>
                <LoginInput onLogin={onLoginHandler} />
                <p>{linkQuestion} <Link to={'/register'}>{linkMessage}</Link></p>
            </section>
        )
    }
    const enRet = returnComponent({heading: 'You must login first to enter this app!', linkQuestion: 'Doesn\'t have an account yet?', linkMessage: 'Register Here'})
    const idRet = returnComponent({heading: 'Login dulu ya!', linkQuestion: 'Belum punya akun?', linkMessage: 'Daftar Di Sini'})

    return localeContext.locale === 'en' ? enRet : idRet
}