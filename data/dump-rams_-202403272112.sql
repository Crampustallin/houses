PGDMP                         |            rams_    14.10 (Homebrew)    14.10 (Homebrew)     $           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            %           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            &           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            '           1262    16670    rams_    DATABASE     P   CREATE DATABASE rams_ WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';
    DROP DATABASE rams_;
             	   mukhammed    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             	   mukhammed    false            (           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                	   mukhammed    false    3            �            1259    16679 	   addresses    TABLE       CREATE TABLE public.addresses (
    address_id integer NOT NULL,
    city character varying(100) NOT NULL,
    street character varying(100) NOT NULL,
    house_number character varying(10) NOT NULL,
    apartment_number character varying(10),
    postal_code character varying(20)
);
    DROP TABLE public.addresses;
       public         heap 	   mukhammed    false    3            �            1259    16678    addresses_address_id_seq    SEQUENCE     �   CREATE SEQUENCE public.addresses_address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.addresses_address_id_seq;
       public       	   mukhammed    false    212    3            )           0    0    addresses_address_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.addresses_address_id_seq OWNED BY public.addresses.address_id;
          public       	   mukhammed    false    211            �            1259    16686 
   properties    TABLE     �   CREATE TABLE public.properties (
    property_id integer NOT NULL,
    property_type_id integer NOT NULL,
    address_id integer NOT NULL,
    price numeric(20,2) NOT NULL,
    rooms integer,
    area numeric(10,2),
    description text
);
    DROP TABLE public.properties;
       public         heap 	   mukhammed    false    3            �            1259    16685    properties_property_id_seq    SEQUENCE     �   CREATE SEQUENCE public.properties_property_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.properties_property_id_seq;
       public       	   mukhammed    false    3    214            *           0    0    properties_property_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.properties_property_id_seq OWNED BY public.properties.property_id;
          public       	   mukhammed    false    213            �            1259    16672    property_types    TABLE     �   CREATE TABLE public.property_types (
    property_type_id integer NOT NULL,
    property_type_name character varying(100) NOT NULL
);
 "   DROP TABLE public.property_types;
       public         heap 	   mukhammed    false    3            �            1259    16671 #   property_types_property_type_id_seq    SEQUENCE     �   CREATE SEQUENCE public.property_types_property_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 :   DROP SEQUENCE public.property_types_property_type_id_seq;
       public       	   mukhammed    false    210    3            +           0    0 #   property_types_property_type_id_seq    SEQUENCE OWNED BY     k   ALTER SEQUENCE public.property_types_property_type_id_seq OWNED BY public.property_types.property_type_id;
          public       	   mukhammed    false    209            �           2604    16682    addresses address_id    DEFAULT     |   ALTER TABLE ONLY public.addresses ALTER COLUMN address_id SET DEFAULT nextval('public.addresses_address_id_seq'::regclass);
 C   ALTER TABLE public.addresses ALTER COLUMN address_id DROP DEFAULT;
       public       	   mukhammed    false    211    212    212            �           2604    16689    properties property_id    DEFAULT     �   ALTER TABLE ONLY public.properties ALTER COLUMN property_id SET DEFAULT nextval('public.properties_property_id_seq'::regclass);
 E   ALTER TABLE public.properties ALTER COLUMN property_id DROP DEFAULT;
       public       	   mukhammed    false    214    213    214            �           2604    16675    property_types property_type_id    DEFAULT     �   ALTER TABLE ONLY public.property_types ALTER COLUMN property_type_id SET DEFAULT nextval('public.property_types_property_type_id_seq'::regclass);
 N   ALTER TABLE public.property_types ALTER COLUMN property_type_id DROP DEFAULT;
       public       	   mukhammed    false    210    209    210                      0    16679 	   addresses 
   TABLE DATA           j   COPY public.addresses (address_id, city, street, house_number, apartment_number, postal_code) FROM stdin;
    public       	   mukhammed    false    212   �!       !          0    16686 
   properties 
   TABLE DATA           p   COPY public.properties (property_id, property_type_id, address_id, price, rooms, area, description) FROM stdin;
    public       	   mukhammed    false    214   �!                 0    16672    property_types 
   TABLE DATA           N   COPY public.property_types (property_type_id, property_type_name) FROM stdin;
    public       	   mukhammed    false    210   	"       ,           0    0    addresses_address_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.addresses_address_id_seq', 3, true);
          public       	   mukhammed    false    211            -           0    0    properties_property_id_seq    SEQUENCE SET     I   SELECT pg_catalog.setval('public.properties_property_id_seq', 1, false);
          public       	   mukhammed    false    213            .           0    0 #   property_types_property_type_id_seq    SEQUENCE SET     R   SELECT pg_catalog.setval('public.property_types_property_type_id_seq', 1, false);
          public       	   mukhammed    false    209            �           2606    16684    addresses addresses_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (address_id);
 B   ALTER TABLE ONLY public.addresses DROP CONSTRAINT addresses_pkey;
       public         	   mukhammed    false    212            �           2606    16693    properties properties_pkey 
   CONSTRAINT     a   ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pkey PRIMARY KEY (property_id);
 D   ALTER TABLE ONLY public.properties DROP CONSTRAINT properties_pkey;
       public         	   mukhammed    false    214            �           2606    16677 "   property_types property_types_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY public.property_types
    ADD CONSTRAINT property_types_pkey PRIMARY KEY (property_type_id);
 L   ALTER TABLE ONLY public.property_types DROP CONSTRAINT property_types_pkey;
       public         	   mukhammed    false    210            �           2606    16699 %   properties properties_address_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.addresses(address_id);
 O   ALTER TABLE ONLY public.properties DROP CONSTRAINT properties_address_id_fkey;
       public       	   mukhammed    false    214    212    3468            �           2606    16694 +   properties properties_property_type_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_property_type_id_fkey FOREIGN KEY (property_type_id) REFERENCES public.property_types(property_type_id);
 U   ALTER TABLE ONLY public.properties DROP CONSTRAINT properties_property_type_id_fkey;
       public       	   mukhammed    false    3466    210    214                  x������ � �      !      x������ � �            x������ � �     