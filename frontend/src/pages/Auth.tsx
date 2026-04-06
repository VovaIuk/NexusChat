import AuthLayout from "../components/auth/AuthLayout";
import RegisterForm from "../components/auth/RegisterForm";

export default function Auth() {
  return (
    <AuthLayout>
      <RegisterForm />
    </AuthLayout>
  );
}
