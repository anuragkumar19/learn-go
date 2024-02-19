'use client'

import { loadHomePageData } from '@/lib/loader'
import { use } from 'react'

const promise = loadHomePageData('')

export const dynamic = 'force-dynamic'

export default function Home() {
    const data = use(promise)

    return (
        <main>
            <ul>
                {data.map((item, i) => (
                    <li>
                        <h1>{item.text}</h1>
                    </li>
                ))}
            </ul>
        </main>
    )
}
