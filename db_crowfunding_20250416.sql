PGDMP                      }            crawfounding    16.6    16.6                 0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    33034    crawfounding    DATABASE     �   CREATE DATABASE crawfounding WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE crawfounding;
                Lenovo    false            �            1259    33054    campaign_images    TABLE       CREATE TABLE public.campaign_images (
    id integer NOT NULL,
    campaign_id integer,
    file_name character varying(255),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    is_primary smallint DEFAULT 0
);
 #   DROP TABLE public.campaign_images;
       public         heap    postgres    false            �            1259    33053    campaign_images_id_seq    SEQUENCE     �   CREATE SEQUENCE public.campaign_images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.campaign_images_id_seq;
       public          postgres    false    220                       0    0    campaign_images_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.campaign_images_id_seq OWNED BY public.campaign_images.id;
          public          postgres    false    219            �            1259    33045 	   campaigns    TABLE     �  CREATE TABLE public.campaigns (
    id integer NOT NULL,
    user_id integer,
    name character varying(255),
    short_description character varying(255),
    description text,
    perks text,
    backer_count integer,
    goal_amount integer,
    current_amount integer,
    slug character varying(255),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.campaigns;
       public         heap    postgres    false            �            1259    33044    campaigns_id_seq    SEQUENCE     �   CREATE SEQUENCE public.campaigns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.campaigns_id_seq;
       public          postgres    false    218                       0    0    campaigns_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.campaigns_id_seq OWNED BY public.campaigns.id;
          public          postgres    false    217            �            1259    33070    transactions    TABLE       CREATE TABLE public.transactions (
    id integer NOT NULL,
    campaign_id integer,
    user_id integer,
    amount integer,
    status character varying(255),
    code character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
     DROP TABLE public.transactions;
       public         heap    postgres    false            �            1259    33069    transactions_id_seq    SEQUENCE     �   CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.transactions_id_seq;
       public          postgres    false    222                       0    0    transactions_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;
          public          postgres    false    221            �            1259    33036    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255),
    occupation character varying(255),
    email character varying(255),
    password_hash character varying(255),
    avatar_file_name character varying(255),
    role character varying(255),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    33035    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    216                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    215            e           2604    33057    campaign_images id    DEFAULT     x   ALTER TABLE ONLY public.campaign_images ALTER COLUMN id SET DEFAULT nextval('public.campaign_images_id_seq'::regclass);
 A   ALTER TABLE public.campaign_images ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    219    220    220            b           2604    33048    campaigns id    DEFAULT     l   ALTER TABLE ONLY public.campaigns ALTER COLUMN id SET DEFAULT nextval('public.campaigns_id_seq'::regclass);
 ;   ALTER TABLE public.campaigns ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            i           2604    33073    transactions id    DEFAULT     r   ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);
 >   ALTER TABLE public.transactions ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    222    222            _           2604    33039    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216                      0    33054    campaign_images 
   TABLE DATA           i   COPY public.campaign_images (id, campaign_id, file_name, created_at, updated_at, is_primary) FROM stdin;
    public          postgres    false    220   �%                 0    33045 	   campaigns 
   TABLE DATA           �   COPY public.campaigns (id, user_id, name, short_description, description, perks, backer_count, goal_amount, current_amount, slug, created_at, updated_at) FROM stdin;
    public          postgres    false    218   &                 0    33070    transactions 
   TABLE DATA           n   COPY public.transactions (id, campaign_id, user_id, amount, status, code, created_at, updated_at) FROM stdin;
    public          postgres    false    222   '                 0    33036    users 
   TABLE DATA           {   COPY public.users (id, name, occupation, email, password_hash, avatar_file_name, role, created_at, updated_at) FROM stdin;
    public          postgres    false    216   $'                  0    0    campaign_images_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.campaign_images_id_seq', 3, true);
          public          postgres    false    219                       0    0    campaigns_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.campaigns_id_seq', 3, true);
          public          postgres    false    217                       0    0    transactions_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);
          public          postgres    false    221                       0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 54, true);
          public          postgres    false    215            o           2606    33059 $   campaign_images campaign_images_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.campaign_images
    ADD CONSTRAINT campaign_images_pkey PRIMARY KEY (id);
 N   ALTER TABLE ONLY public.campaign_images DROP CONSTRAINT campaign_images_pkey;
       public            postgres    false    220            m           2606    33052    campaigns campaigns_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.campaigns DROP CONSTRAINT campaigns_pkey;
       public            postgres    false    218            q           2606    33077    transactions transactions_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.transactions DROP CONSTRAINT transactions_pkey;
       public            postgres    false    222            k           2606    33043    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216               ^   x�}�1
�0�:�
?�{1�%o��F���]��z`b`���<���{���ʼ�b�U���A�?6x�Tw��A����RQX�(����a-�         �   x�}�A�� E��)� j��e6�@)cB,���>Ĩ����Ճ�k8�u$�9�����A��s�hoO���2��3Ƒ�PSJ���ho�`q$��ҹ`4~�@q�pF=�B�vV�U"_%$�*�\TțV��BUUUW�DL��	F��zxa4%��	q��0,��u����T��k�7�x�˲��E��H�fS+<,7�}궉"�n��q��/��oNu����K��N�W��uh�d            x������ � �         W  x���AO�@�ϓ_�C��c����D�Qڥ�zA+�i2I�6�4-����@�����{�g+{�ح��ݞ���zYVv��]%�h���ʂwM>��/;^�L��魻�������3{Z�y9����R�⟙��]���B�� TB�(
#�Z���D���5i@�^o�ovec��/zړ]]~�y�ʖ���5E�(�Q���s�HI����/IT�j�ӊ���I�#Cht�$���bq=H]�.��,�G����l��cs��ǻ�"*�wϱ��ޑ[.}�p,6w��P��NT��Ф83Q,�Ȩ~,�3Fd$b�����u���d������a/���     