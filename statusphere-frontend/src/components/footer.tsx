import Link from 'next/link'

import Image from 'next/image'
import { Container } from './mainsite/Container'
import {NavLink} from './mainsite/NavLink'
import icon from '../../public/static/images/logos/icon.png'
import React from "react";

export function Footer() {
    return (
        <footer className={"mt-8"}>
            <Container>
                <div className="py-16">
                    <Image className={"mx-auto h-10 w-auto"} src={icon} alt="Metoro" width={0} height={0}/>
                    <nav className="mt-10 text-sm" aria-label="quick links">
                        <div className="-my-1 flex justify-center gap-x-6">
                            <NavLink href="https://metoro.io#features">Features</NavLink>
                            <NavLink href="https://metoro.io#use-cases">Use Cases</NavLink>
                            <NavLink href="https://metoro.io#pricing">Pricing</NavLink>
                            <NavLink href="https://metoro.io#faqs">FAQs</NavLink>
                            {/*<NavLink href="#demo">Live Demo</NavLink>*/}
                            <NavLink href="https://docs.metoro.io">Docs</NavLink>
                            <NavLink href="https://metoro.io/blog">Blog</NavLink>
                        </div>
                    </nav>
                </div>
                <div
                    className="flex flex-col items-center border-t border-slate-400/10 py-10 sm:flex-row-reverse sm:justify-between">
                    <div className="flex gap-x-6">
                        <Link href="https://www.linkedin.com/company/metoroai" className="group" aria-label="Metoro on X">
                            <svg
                                className="h-6 w-6 fill-slate-500 group-hover:fill-slate-700"
                                aria-hidden="true"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    d="M19 0h-14c-2.761 0-5 2.239-5 5v14c0 2.761 2.239 5 5 5h14c2.762 0 5-2.239 5-5v-14c0-2.761-2.238-5-5-5zm-11 19h-3v-11h3v11zm-1.5-12.268c-.966 0-1.75-.79-1.75-1.764s.784-1.764 1.75-1.764 1.75.79 1.75 1.764-.783 1.764-1.75 1.764zm13.5 12.268h-3v-5.604c0-3.368-4-3.113-4 0v5.604h-3v-11h3v1.765c1.396-2.586 7-2.777 7 2.476v6.759z"/>
                            </svg>
                        </Link>
                        <Link href="https://twitter.com/metoro_ai" className="group"
                              aria-label="Metoro on X">
                            <svg
                                className="h-6 w-6 fill-slate-500 group-hover:fill-slate-700"
                                aria-hidden="true"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    d="M13.3174 10.7749L19.1457 4H17.7646L12.7039 9.88256L8.66193 4H4L10.1122 12.8955L4 20H5.38119L10.7254 13.7878L14.994 20H19.656L13.3171 10.7749H13.3174ZM11.4257 12.9738L10.8064 12.0881L5.87886 5.03974H8.00029L11.9769 10.728L12.5962 11.6137L17.7652 19.0075H15.6438L11.4257 12.9742V12.9738Z"/>
                            </svg>
                        </Link>
                        {/*<Link href="#" className="group" aria-label="Metoro on GitHub">*/}
                        {/*    <svg*/}
                        {/*        className="h-6 w-6 fill-slate-500 group-hover:fill-slate-700"*/}
                        {/*        aria-hidden="true"*/}
                        {/*        viewBox="0 0 24 24"*/}
                        {/*    >*/}
                        {/*        <path*/}
                        {/*            d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0 1 12 6.844a9.59 9.59 0 0 1 2.504.337c1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.02 10.02 0 0 0 22 12.017C22 6.484 17.522 2 12 2Z"/>*/}
                        {/*    </svg>*/}
                        {/*</Link>*/}
                    </div>
                    <p className="mt-6 text-sm text-slate-500 sm:mt-0">
                        Copyright &copy; {new Date().getFullYear()} Metoro Inc. All rights
                        reserved.
                    </p>
                </div>
            </Container>
        </footer>
    )
}