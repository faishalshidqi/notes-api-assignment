import {showFormattedDate} from '../utils/data.ts'
import {Link} from 'react-router-dom'
import DeleteButton from './DeleteButton.tsx'
import LocaleContext from '../contexts/LocaleContext.ts'
import { useContext } from 'react'

export default function NoteItem({id, title, body, created_at, onDelete}: {id: string, title: string, body: string, created_at: string, onDelete: (id: string) => void}) {
    const localeContext = useContext(LocaleContext)
    return (
        <article className='note-item'>
            <h3 className='note-item__title'>
                <Link to={`/notes/app/${id}`}>{title}</Link>
            </h3>
            <p className='note-item__createdAt'>{showFormattedDate(created_at, localeContext.locale === 'en' ? 'en-US' : 'id-ID')}</p>
            <p className='note-item__body'>{body}</p>
            <div className='note-item__action'>
                <DeleteButton onDelete={() => onDelete(id)} id={id}></DeleteButton>
            </div>
        </article>
    )
}