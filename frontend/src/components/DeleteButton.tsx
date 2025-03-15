import {FiTrash} from 'react-icons/fi'

export default function DeleteButton({id, onDelete}: {id: string, onDelete: (id: string) => void}) {
    return (
        <button type='button' className='action' title='Delete' onClick={() => onDelete(id)}><FiTrash/></button>
    )
}