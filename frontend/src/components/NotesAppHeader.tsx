import Navigation from './Navigation.tsx'
import {Link} from 'react-router-dom'
import {useContext} from 'react'
import LocaleContext from '../contexts/LocaleContext.ts'

export default function NotesAppHeader({onLogout}: {onLogout: () => void}) {
    const localeContext = useContext(LocaleContext)
    return (
        <header>
            <h1>
                <Link to='/'>{localeContext.locale === 'en' ? 'Notes App' : 'Aplikasi Catatan'}</Link>
            </h1>
            <Navigation onLogout={onLogout} />
        </header>
    )
}