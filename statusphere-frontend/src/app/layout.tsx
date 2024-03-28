import type {Metadata} from "next";
import {Inter, Lexend} from "next/font/google";
import '@/index.css'
import {clsx} from "clsx";
import {PHProvider} from "@/app/providers";
import PostHogPageView from "@/app/PostHogPageView";

const inter = Inter({
    subsets: ['latin'],
    display: 'swap',
    variable: '--font-inter',
})

const lexend = Lexend({
    subsets: ['latin'],
    display: 'swap',
    variable: '--font-lexend',
})

export const metadata: Metadata = {
    title: "Statusphere",
    description: "An status page aggregator",
};

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html
            lang="en"
            className={clsx(
                'h-full scroll-smooth bg-white antialiased',
                inter.variable,
                lexend.variable,
            )}
        >
        <PHProvider>
            <body className={"flex h-full flex-col"}>
            <PostHogPageView/>
            {children}
            </body>
        </PHProvider>

        </html>
    );
}
