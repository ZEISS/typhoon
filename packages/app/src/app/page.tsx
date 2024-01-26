import { Metadata } from 'next'
import DefaultLayout from '@/components/default-layout'
import { PropsWithChildren } from 'react'

export const metadata: Metadata = {
  title: 'Typhoon',
  description: ''
}

export interface RootProps extends PropsWithChildren { }

export default async function Root({ children }: Readonly<RootProps>) {
  return (
    <>
      <DefaultLayout>{children}</DefaultLayout>
    </>
  )
}
