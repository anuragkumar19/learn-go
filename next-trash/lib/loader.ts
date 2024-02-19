export function loadHomePageData(_: string) {
    return new Promise<{ text: string }[]>((res) => {
        setTimeout(() => {
            res([
                {
                    text: 'Buy pizza',
                },
                {
                    text: 'Buy Vercel',
                },
                {
                    text: 'Buy Next.js',
                },
                {
                    text: 'Buy Milk',
                },
                {
                    text: 'Buy AWS',
                },
            ])
        }, 1000)
    })
}
