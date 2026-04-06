import type { ReactNode } from "react";

interface AuthCardProps {
  title: string;
  subtitle: string;
  children: ReactNode;
}

export default function AuthCard({ title, subtitle, children }: AuthCardProps) {
  return (
    <section className="auth-card">
      <h1 className="auth-title">{title}</h1>
      <p className="auth-subtitle">{subtitle}</p>
      {children}
    </section>
  );
}
