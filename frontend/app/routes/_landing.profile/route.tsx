import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "My Profile" },
    {
      name: "description",
      content: "This is the user profile page of the driver app",
    },
  ];
};

export default function Route() {
  return <div>Landing Profile</div>;
}
