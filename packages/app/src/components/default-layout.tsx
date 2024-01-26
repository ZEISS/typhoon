import '@/styles/globals.css'
import { type ReactNode } from 'react'
import { ThemeToggle } from './theme-toggle'
import { UserNav } from './user-nav'

interface DefaultLayoutProps {
  children: ReactNode
}

export default async function DefaultLayout({
  children
}: Readonly<DefaultLayoutProps>) {
  return (
    <>
      <div className="flex-col">
        <div className="border-b">
          <div className="flex h-16 items-center px-4">
            <div className="ml-auto flex items-center space-x-4">
              <ThemeToggle />
              <UserNav />
            </div>
          </div>
        </div>
        <div className="flex-1">{children}</div>
      </div>
    </>
  )
}
