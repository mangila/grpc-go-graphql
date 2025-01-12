import {NextRequest, NextResponse} from 'next/server';

export function middleware(request: NextRequest): NextResponse {
    console.log('Request URL:', request.url);
    return NextResponse.next();
}
