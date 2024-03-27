import { useRouter } from 'next/router'

export default function StatusPage() {
    const router = useRouter()
    return <p>Post: {router.query.company}</p>
}