import {FaMoon, FaSun} from 'react-icons/fa'
import {FiLogOut} from 'react-icons/fi'
import LocaleContext from '../contexts/LocaleContext.ts'
import {useContext} from 'react'
import {MdGTranslate} from 'react-icons/md'
import UserAuthContext from '../contexts/UserAuthContext.ts'
import ThemeContext from '../contexts/ThemeContext.ts'

export default function Navigation({onLogout}: {onLogout: () => void}) {
    const userContext = useContext(UserAuthContext)
    const themeContext = useContext(ThemeContext)
    const locale = useContext(LocaleContext)
    const authed = (
        <nav className='navigation'>
            <ul>
                <li>
                    <a onClick={locale.toggleLocale}><MdGTranslate/></a>
                </li>
                <li>
                    <a onClick={themeContext.toggleTheme}>{themeContext.theme === 'light' ? <FaSun/> : <FaMoon/>}</a>
                </li>
                <li>
                    <a onClick={onLogout}><FiLogOut/></a>
                </li>
            </ul>
        </nav>
    )
    const not_authed = (
        <nav className='navigation'>
            <ul>
                <li>
                    <a onClick={locale.toggleLocale}><MdGTranslate/></a>
                </li>
                <li>
                    <a onClick={themeContext.toggleTheme}>{themeContext.theme === 'light' ? <FaSun/> : <FaMoon/>}</a>
                </li>
            </ul>
        </nav>
    )
    return userContext.id === '' ? not_authed : authed
}