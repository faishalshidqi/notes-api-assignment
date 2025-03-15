import NoteInput from '../components/NoteInput.tsx'
import {addNote} from '../utils/data.ts'
import {useNavigate} from 'react-router-dom'
import {useContext} from 'react'
import LocaleContext from '../contexts/LocaleContext.ts'

export default function AddNotePage() {
    const navigate = useNavigate()
    const localeContext = useContext(LocaleContext)
    const onAddNoteHandler = ({title, body}: {title: string, body: string}): void => {
        addNote({title, body})
        navigate('/')
    }
    return (
        <div className='add-new-page'>
            <h2>{localeContext.locale === 'en' ? 'Add Note' : 'Tambah Catatan'}</h2>
            <NoteInput addNote={onAddNoteHandler} />
        </div>
    )
}