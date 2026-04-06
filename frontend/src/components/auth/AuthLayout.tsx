import type { ReactNode } from "react";
import "../chat/theme.css";
import "./auth.css";

interface AuthLayoutProps {
  children: ReactNode;
}

export default function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <div className="theme-provider" data-theme="light">
      <main className="auth-page">{children}</main>
    </div>
  );
}
