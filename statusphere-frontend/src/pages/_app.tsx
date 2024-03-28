import '@/index.css'
import {AppProps} from "next/app";
import Layout from "@/components/Layout";
import posthog from "posthog-js";
import {useRouter} from "next/router";
import {useEffect} from "react";
import {PostHogProvider} from "posthog-js/react";


// Check that PostHog is client-side (used to handle Next.js SSR)
if (typeof window !== 'undefined') {
    posthog.init(process.env.NEXT_PUBLIC_POSTHOG_KEY!, {
        api_host: process.env.NEXT_PUBLIC_POSTHOG_HOST || 'https://app.posthog.com',
        // Enable debug mode in development
        loaded: (posthog) => {
            if (process.env.NODE_ENV === 'development') posthog.debug()
        }
    })
}

export default function MyApp({Component, pageProps}: AppProps) {
    const router = useRouter()

    useEffect(() => {
        // Track page views
        const handleRouteChange = () => posthog?.capture('$pageview')
        router.events.on('routeChangeComplete', handleRouteChange)

        return () => {
            router.events.off('routeChangeComplete', handleRouteChange)
        }
    }, [])

    return <PostHogProvider>
        <Layout>
            <Component {...pageProps} />
        </Layout>
    </PostHogProvider>

}