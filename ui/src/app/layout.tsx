import React from "react";
import type {Metadata} from "next";
import {AppRouterCacheProvider} from "@mui/material-nextjs/v13-appRouter";
import {ThemeProvider} from "@mui/system";
import theme from '../theme';
import {CssBaseline} from "@mui/material";

export const metadata: Metadata = {
    title: "Mangila@Github",
    description: "Order, products and customers",
};

export default function RootLayout({children}: Readonly<{ children: React.ReactNode }>) {
    return (
        <html lang="en">
        <body>
        <AppRouterCacheProvider options={{enableCssLayer: true}}>
            <CssBaseline/>
            <ThemeProvider theme={theme}>
                {children}
            </ThemeProvider>
        </AppRouterCacheProvider>
        </body>
        </html>
    );
}
