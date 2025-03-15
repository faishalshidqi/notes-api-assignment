import RegisterInput from '../components/RegisterInput'
import {Link, useNavigate} from 'react-router-dom'
import {register} from '../utils/data.ts'
import {useContext} from 'react'
import LocaleContext from '../contexts/LocaleContext.ts'

export default function RegisterPage() {
    const localeContext = useContext(LocaleContext)
    const navigate = useNavigate()
    async function onRegisterHandler({fullname, username, password}: {fullname: string, username: string, password: string}) {
        const {error} = await register({fullname, username, password})
        if (!error) navigate('/')
    }

    function returnComponent({heading, linkQuestion, linkMessage}: {heading: string, linkQuestion: string, linkMessage: string}) {
        return (
            <section className={'register-page'}>
                <h2>{heading}</h2>
                <RegisterInput onRegister={onRegisterHandler} />
                <p>{linkQuestion} <Link to={'/'}>{linkMessage}</Link></p>
            </section>
        )
    }

    const enRet = returnComponent({heading: 'Fill the form to register an account!', linkQuestion: 'Already have an account?', linkMessage: 'Login Here'})
    const idRet = returnComponent({heading: 'Isi form berikut untuk mendaftarkan akun!', linkQuestion: 'Sudah punya akun?', linkMessage: 'Login Di Sini'})

    return localeContext.locale === 'en' ? enRet : idRet
}